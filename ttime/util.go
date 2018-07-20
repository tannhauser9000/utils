package ttime;

import "fmt";
import "time";

func GetTimestamp() (int64) {
  return time.Now().Unix();
}

func GetDate() (string) {
  return fmt.Sprintf("%s", time.Now().UTC().Format("2006-01-02T15:04:05Z"));
}

func GetTimestampDate() (int64, string) {
  now := time.Now();
  return now.Unix(), fmt.Sprintf("%s", now.UTC().Format("2006-01-02T15:04:05Z"));
}
