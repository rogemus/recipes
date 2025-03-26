export interface Response<Body = unknown, Error = unknown> {
  data?: Body;
  error?: Error;
}
