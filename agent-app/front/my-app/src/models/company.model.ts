export class Company {
  id: string = '';
  companyName: string = '';
  ownerId: string = '';
  description: string = '';
  info: Info = new Info();
  openJobs: Job[] = [];
  comments: Comment[] = [];
}

class Info {
  street: string = '';
	city: string = '';
	country: string = '';
	phoneNumber: string = '';
	email: string = '';
}

export class Job {
  id: string = '';
  jobName: string = '';
  jobDescription: string = ''; 
}

export class Comment {
  id: string = '';
	commentOwnerId: string = '';
	comment: string = '';
  grade: number = 1;
}

export class CompaniesResponse {
  companies: Company[] = [];
}