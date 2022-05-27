export class User {
  id: String = '';
  username: String = '';
  firstName: String = '';
  lastName: String = '';
  email: String = '';
  phoneNumber: String = '';
  biography: String = '';
  birthDate: Date = new Date();
  role: String = '';
  image: String ='';
  // TODO:
  // repeated EducationItem education = 8;
  // repeated WorkExpirienceItem workExpirience = 9;
  // repeated SkillItem skills = 10;
  // repeated Interests interests = 11;
  // repeated string blockedUsers = 12;
}

export class UserResponse {
  user: User = new User();
}

export class ManyUserResponse {
  users: User[] = [];
}