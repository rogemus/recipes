export interface Response<Body = {}> {
  data?: Body;
  error?: {
    msg: string;
  };
}
