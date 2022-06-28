import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ChangeUserComponent } from './components/change-user/change-user.component';
import { CompanyProfileComponent } from './components/company-profile/company-profile.component';
import { CreateCommentComponent } from './components/create-comment/create-comment.component';
import { CreateCompanyComponent } from './components/create-company/create-company.component';
import { CreateJobComponent } from './components/create-job/create-job.component';
import { EnableCompanyComponent } from './components/enable-company/enable-company.component';
import { LoginComponent } from './components/login/login.component';
import { MyCompaniesComponent } from './components/my-companies/my-companies.component';
import { RegistrationComponent } from './components/registration/registration.component';
import { UserProfileComponent } from './components/user-profile/user-profile.component';

const routes: Routes = [
  {path: 'login', component: LoginComponent},
  {path: 'register', component: RegistrationComponent},

  {path: 'userProfile', component: UserProfileComponent},
  {path: 'userProfile/createCompany', component: CreateCompanyComponent},
  {path: 'userProfile/changeUser', component: ChangeUserComponent},
  {path: 'userProfile/myCompanies', component: MyCompaniesComponent},

  {path: 'enableCompany', component: EnableCompanyComponent},
  {path: 'companyProfile', component: CompanyProfileComponent},
  {path: 'createJob', component: CreateJobComponent},
  {path: 'createComment', component: CreateCommentComponent},

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
