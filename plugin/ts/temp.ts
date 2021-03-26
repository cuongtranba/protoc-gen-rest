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
