package numbers;

type Error string;

/* for constant error */
func (e Error) Error() (string) {
  return string(e);
}

const ErrInputBufferLength = Error("insufficient buffer size of the input buffer");
const ErrOutputBufferLength = Error("insufficient buffer size of the output buffer");
