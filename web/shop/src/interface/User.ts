interface User {
  firstname: string;
  lastname: string;
  email: string;
  phone: string;
  iss: string;
  sub: string;
  aud: [string];
  exp: number;
}

export default User;
