export class User {
  id: String = '';
  username: String = '';
  firstName: String = '';
  lastName: String = '';
  email: String = '';
  phoneNumber: String = '';
  birthDate: Date = new Date();
  role: String = '';
  profileImage: String = '';
  password: String = '';
}

export class UserResponse {
  user: User = new User();
}