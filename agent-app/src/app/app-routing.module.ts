import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { MainPageComponent } from './components/main-page/main-page.component';
import { ProfileComponent } from './components/profileComponents/profile/profile.component';
import { RegistrationComponent } from './components/registration/registration.component';

const routes: Routes = [
  {path: 'login', component: LoginComponent},
  {path: 'register', component: RegistrationComponent},
  {path: 'main', component: MainPageComponent},
  {path: 'profile', component: ProfileComponent},

  // ovde se dodaje komponenta i moguce ju je direktno pogoditi kroz url
  // primer recimo da postoji komponenta koja jos uvek nije povezana vec je treba samo testirati:
  //
  // {path: 'createPost', component: CreatePostComponent},
  //
  // putanja do nje je http://localhost:4200/createPost

  // sto se tice registracije ona postoji na pritisak dugmeta za registraciju
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
