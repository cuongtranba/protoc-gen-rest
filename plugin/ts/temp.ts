import axios, { AxiosResponse, AxiosRequestConfig } from "axios";
import * as student from "./student";
export interface Scalars {
  double: number;
  float: number;
  int32: number;
  int64: number;
  uint32: number;
  uint64: number;
  sint32: number;
  sint64: number;
  fixed32: number;
  fixed64: number;
  sfixed32: number;
  sfixed64: number;
  bool: boolean;
  string: string;
  bytes: number[];
  name?: string;
  User: User;
  Users: User[];
  Student: student.Student[];
}
export interface User {
  Name: string;
  Age: string;
}

export const getUser = (config?: AxiosRequestConfig) => async (
  request: User
): Promise<AxiosResponse<Scalars>> => {
  const result = await axios.post("/getuser", request, config);
  return result.data;
};

export const UserService = (config?: AxiosRequestConfig) => {
  getUser: getUser(config);
};
