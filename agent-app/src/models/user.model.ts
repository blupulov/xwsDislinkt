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

// ZBOG FORMATA U KOJEM SERVER VRACA OVAJ PODATAK "user": {"id":"...", ....}
export class UserResponse {
  user: User = new User();
}

export class ManyUserResponse {
  users: User[] = [];
}