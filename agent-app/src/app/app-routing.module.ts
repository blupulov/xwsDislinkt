import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { MainPageComponent } from './components/main-page/main-page.component';
import { ChangeInfoComponent } from './components/profileComponents/change-info/change-info.component';
import { FollowingUsersComponent } from './components/profileComponents/following-users/following-users.component';
import { ProfileComponent } from './components/profileComponents/profile/profile.component';
import { UserFollowersComponent } from './components/profileComponents/user-followers/user-followers.component';
import { UserPostsComponent } from './components/profileComponents/user-posts/user-posts.component';
import { RegistrationComponent } from './components/registration/registration.component';

const routes: Routes = [
  {path: 'login', component: LoginComponent},
  {path: 'register', component: RegistrationComponent},
  {path: 'main', component: MainPageComponent},

  //KOMPONENTE KOJIMA SE PRISTUPA SA PROFILA ULOGOVANOG KORISNIKA
  {path: 'profile', component: ProfileComponent},
  {path: 'profile/changeInfo', component: ChangeInfoComponent},
  {path: 'profile/myFollowers', component: UserFollowersComponent},
  {path: 'profile/following', component: FollowingUsersComponent},
  {path: 'profile/myPosts', component: UserPostsComponent},

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
