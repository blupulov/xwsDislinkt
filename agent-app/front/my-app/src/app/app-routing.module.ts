import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ChangeUserComponent } from './components/change-user/change-user.component';
import { CreateCompanyComponent } from './components/create-company/create-company.component';
import { EnableCompanyComponent } from './components/enable-company/enable-company.component';
import { LoginComponent } from './components/login/login.component';
import { RegistrationComponent } from './components/registration/registration.component';
import { UserProfileComponent } from './components/user-profile/user-profile.component';

const routes: Routes = [
  {path: 'login', component: LoginComponent},
  {path: 'register', component: RegistrationComponent},
  {path: 'userProfile', component: UserProfileComponent},
  {path: 'userProfile/createCompany', component: CreateCompanyComponent},
  {path: 'userProfile/changeUser', component: ChangeUserComponent},
  {path: 'enableCompany', component: EnableCompanyComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
