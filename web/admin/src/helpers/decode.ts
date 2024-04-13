import User from "@/interface/User";
import { jwtDecode } from "jwt-decode";

export function decodeToken(token: string): User | null {
  try {
    const decoded = jwtDecode(token);
    return decoded as User;
  } catch (error) {
    return null;
  }
}
