/*
formatter
request format checker
*/

package formatter;

import "encoding/json";
import "os";
import "strings";

import "reflect";
import "log";

type _rule_set_st struct {
  Rules []*_rule_st `json:"rules"`;
}

type _rule_st struct {
  Key string `json:"key"`;
  Type string `json:"type"`;
  Require bool `json:"require"`;
  Default interface{} `json:"default"`;
  Limit int `json:"limit"`;
}

type RuleSet struct {
  debug bool;
  rules map[string]*Rule;
}

type Rule struct {
  key string;
  type_value string;
  require bool;
  default_value interface{};
  limit int;
}

type Request struct {
  items map[string]interface{};
  valid bool;
}

type Error string;

func (e Error) Error() (string) {
  return string(e);
}

const _limit_size_conf = int64(1048576);    // limit the conf file size, preventing reading an extremely large file

const ErrOverSize = Error("configuration oversized, the limitation is 1MB");
const ErrMissingField = Error("missing required field");
const ErrTypeMismatch = Error("type mismatch for field");
const ErrTypeAssert = Error("type assert failed, field type mismatch or empty field");
const ErrFieldOversized = Error("field is oversized");

const NilString = "n/a";

func GetRuleSet(path string, debug bool) (*RuleSet, error) {
  fp, err := os.Open(path);
  defer fp.Close();
  if err != nil {
    return nil, err;
  }
  var fi os.FileInfo;
  fi , err = fp.Stat();
  if err != nil {
    return nil, err;
  }
  if fi.Size() > _limit_size_conf {
    return nil, ErrOverSize;
  }
  decoder := json.NewDecoder(fp);
  rs := &_rule_set_st {};
  err = decoder.Decode(rs);
  if err != nil {
    return nil, err;
  }
  rule_set := &RuleSet {
    debug: debug,
    rules: make(map[string]*Rule),
  };
  for i := 0; i < len((*rs).Rules); i++ {
    (*rule_set).rules[(*(*rs).Rules[i]).Key] = &Rule {
      key: (*(*rs).Rules[i]).Key,
      type_value: (*(*rs).Rules[i]).Type,
      require: (*(*rs).Rules[i]).Require,
      default_value: (*(*rs).Rules[i]).Default,
      limit: (*(*rs).Rules[i]).Limit,
    };
    (*rs).Rules[i] = nil;
  }
  return rule_set, nil;
}

func (rs *RuleSet) CheckRequest(r string, req *Request) (error, string) {
  decoder := json.NewDecoder(strings.NewReader(r));
  request := make(map[string]interface{});
  err := decoder.Decode(&request);
  if err != nil {
    return err, NilString;
  }
  log.Printf("request string: %s\n", r);
  log.Printf("incoming request: %v\n", request);
  if req == nil {
    req = &Request {
      items: make(map[string]interface{}),
      valid: false,
    };
  }
  if (*req).items == nil {
    (*req).items = make(map[string]interface{});
    (*req).valid = false;
  }
  for k, v := range (*rs).rules {
    content, ok := request[k];
    if (*v).require && !ok && (*v).default_value == nil {
      return ErrMissingField, k;
    }
    type_handled := false;
    type_ok := false;
    oversized := false;

    // handle require empty field with default values
    if (*v).require && !ok && !type_handled && (*v).type_value == "int" && (*v).limit > 0 {
      value := float64(0);
      value, type_ok = (*v).default_value.(float64);
      oversized = int(value) > (*v).limit;
      type_handled = true;
    }
    if (*v).require && !ok && !type_handled && (*v).type_value == "int" {
      _, type_ok = (*v).default_value.(float64);
      type_handled = true;
    }
    if (*v).require && !ok && !type_handled && (*v).type_value == "string" && (*v).limit > 0 {
      value := "";
      value, type_ok = (*v).default_value.(string);
      oversized = len(value) > (*v).limit;
      type_handled = true;
    }
    if (*v).require && !ok && !type_handled && (*v).type_value == "base64" && (*v).limit > 0 {
      value := "";
      value, type_ok = (*v).default_value.(string);
      oversized = (len(value) / 4 * 3)  > (*v).limit;
      type_handled = true;
    }
    if (*v).require && !ok && !type_handled && ((*v).type_value == "string" || (*v).type_value == "base64") {
      _, type_ok = (*v).default_value.(string);
      type_handled = true;
    }
    if (*v).require && !ok && !type_handled && (*v).type_value == "bool" {
      _, type_ok = (*v).default_value.(bool);
      type_handled = true;
    }
    if (*v).require && !ok && !type_handled && (*v).type_value == "dict" {
      _, type_ok = (*v).default_value.([]interface{});
    }
    if (*v).require && !ok && !type_handled && (*v).type_value == "dict" && !type_ok {
      type_handled = true;
    }
    if (*v).require && !ok && !type_handled && (*v).type_value == "dict" && type_ok {
      type_dict_ok := true;
      for _, item := range (*v).default_value.([]interface{}) {
        d, type_this := item.(map[string]interface{});
        for _, this_value := range d {
          _, type_value_ok := this_value.(string);
          type_this = type_this && type_value_ok;
        }
        type_dict_ok = type_dict_ok && type_this;
      }
      type_ok = type_ok && type_dict_ok;
    }
    if (*v).require && !ok && !type_handled && (*v).type_value == "dict" && type_ok && (*v).limit > 0 {
      this_size := 0;
      for _, item := range (*v).default_value.([]interface{}) {
        d := item.(map[string]interface{});
        for key, this_value := range d {
          this_size += len(key) + len(this_value.(string));
        }
      }
      oversized = this_size > (*v).limit;
      type_handled = true;
    }
    if (*v).require && !ok && !type_handled && (*v).type_value == "dict" && type_ok {
      type_handled = true;
    }
    if (*v).require && !ok && (*rs).debug {
      log.Printf("[CheckRequest] %v/%v item (%s, %v, %v, %s, %v): %v\n", (*v).require, ok, (*v).type_value, reflect.TypeOf((*v).default_value), type_ok, k, (*v).default_value, request);
    }
    if !ok && !(*v).require {
      type_ok = true;
    }

    // handle non-empty fields
    if ok && !type_handled && (*v).type_value == "int" && (*v).limit > 0 {
      value := float64(0);
      value, type_ok = content.(float64);
      oversized = int(value) > (*v).limit;
      type_handled = true;
    }
    if ok && !type_handled && (*v).type_value == "int" {
      _, type_ok = content.(float64);
      type_handled = true;
    }
    if ok && !type_handled && (*v).type_value == "string" && (*v).limit > 0 {
      value := "";
      value, type_ok = content.(string);
      oversized = len(value) > (*v).limit;
      type_handled = true;
    }
    if ok && !type_handled && (*v).type_value == "base64" && (*v).limit > 0 {
      value := "";
      value, type_ok = content.(string);
      oversized = (len(value) / 4 * 3)  > (*v).limit;
      type_handled = true;
    }
    if ok && !type_handled && ((*v).type_value == "string" || (*v).type_value == "base64") {
      _, type_ok = content.(string);
      type_handled = true;
    }
    if ok && !type_handled && (*v).type_value == "bool" {
      _, type_ok = content.(bool);
      type_handled = true;
    }
    if ok && !type_handled && (*v).type_value == "dict" {
      _, type_ok = content.([]interface{});
    }
    if ok && !type_handled && (*v).type_value == "dict" && !type_ok {
      type_handled = true;
    }
    if ok && !type_handled && (*v).type_value == "dict" && type_ok {
      type_dict_ok := true;
      for _, item := range content.([]interface{}) {
        d, type_this := item.(map[string]interface{});
        for _, this_value := range d {
          _, type_value_ok := this_value.(string);
          type_this = type_this && type_value_ok;
        }
        type_dict_ok = type_dict_ok && type_this;
      }
      type_ok = type_ok && type_dict_ok;
    }
    if ok && !type_handled && (*v).type_value == "dict" && type_ok && (*v).limit > 0 {
      this_size := 0;
      for _, item := range content.([]interface{}) {
        d := item.(map[string]interface{});
        for key, this_value := range d {
          this_size += len(key) + len(this_value.(string));
        }
      }
      oversized = this_size > (*v).limit;
      type_handled = true;
    }
    if ok && !type_handled && (*v).type_value == "dict" && type_ok {
      type_handled = true;
    }
    if (*rs).debug {
      log.Printf("[CheckRequest] %v/%v item (%s, %v, %v, %s, %v): %v\n", (*v).require, ok, (*v).type_value, reflect.TypeOf(content), type_ok, k, content, request);
    }
    if !type_ok {
      return ErrTypeMismatch, k;
    }
    if oversized && (*v).type_value != "int" {
      return ErrFieldOversized, k;
    }
    if ok && !oversized {
      (*req).items[k] = content;
    }
    if !ok && (*v).require && !oversized {
      (*req).items[k] = (*v).default_value;
    }
  }
  (*req).valid = true;
  return nil, NilString;
}

func (rs *RuleSet) Reset(r *Request) {
  if r == nil {
    r = &Request {
      items: make(map[string]interface{}),
      valid: false,
    };
    return;
  }
  for k, v := range (*r).items {
    if (*(*rs).rules[k]).type_value == "dict" {
      for m, _ := range v.(map[string]interface{}) {
        delete(v.(map[string]interface{}), m);
      }
    } else {
      delete ((*r).items, k);
    }
  }
}

func (r *Request) IsValid() (bool) {
  return (*r).valid;
}

func (r *Request) GetStringItem(item string) (string, error) {
  result, ok := (*r).items[item].(string);
  if !ok {
    return "", ErrTypeAssert;
  }
  return result, nil;
}

func (r *Request) GetIntItem(item string) (int, error) {
  result, ok := (*r).items[item].(float64);
  if !ok {
    return -1, ErrTypeAssert;
  }
  return int(result), nil;
}

func (r *Request) GetBoolItem(item string) (bool, error) {
  result, ok := (*r).items[item].(bool);
  if !ok {
    return false, ErrTypeAssert;
  }
  return result, nil;
}

func (r *Request) GetDictItem(item string) ([]interface{}, error) {
  result, ok := (*r).items[item].([]interface{});
  if !ok {
    return nil, ErrTypeAssert;
  }
  return result, nil;
}

