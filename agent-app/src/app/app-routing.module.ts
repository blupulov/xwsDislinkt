import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { MainPageComponent } from './components/main-page/main-page.component';
import { PostFansComponent } from './components/post-fans/post-fans.component';
import { AddEducationComponent } from './components/profileComponents/add-education/add-education.component';
import { AddExpirienceComponent } from './components/profileComponents/add-expirience/add-expirience.component';
import { AddInterestComponent } from './components/profileComponents/add-interest/add-interest.component';
import { AddSkillComponent } from './components/profileComponents/add-skill/add-skill.component';
import { ChangeInfoComponent } from './components/profileComponents/change-info/change-info.component';
import { CreatePostComponent } from './components/profileComponents/create-post/create-post.component';
import { FollowingUsersComponent } from './components/profileComponents/following-users/following-users.component';
import { ProfileComponent } from './components/profileComponents/profile/profile.component';
import { UserFollowersComponent } from './components/profileComponents/user-followers/user-followers.component';
import { UserPostsComponent } from './components/profileComponents/user-posts/user-posts.component';
import { RegistrationComponent } from './components/registration/registration.component';
import { SelectedUserFollowersComponent } from './components/selected-user-followers/selected-user-followers.component';
import { SelectedUserPostsComponent } from './components/selected-user-posts/selected-user-posts.component';
import { SelectedUserProfileComponent } from './components/selected-user-profile/selected-user-profile.component';

const routes: Routes = [
  {path: 'login', component: LoginComponent},
  {path: 'register', component: RegistrationComponent},
  {path: 'main', component: MainPageComponent},

  //KOMPONENTE KOJIMA SE PRISTUPA SA PROFILA ULOGOVANOG KORISNIKA
  {path: 'profile', component: ProfileComponent},
  {path: 'profile/myFollowers', component: UserFollowersComponent},
  {path: 'profile/following', component: FollowingUsersComponent},
  {path: 'profile/myPosts', component: UserPostsComponent},
  {path: 'profile/createPost', component: CreatePostComponent},
  {path: 'profile/changeInfo', component: ChangeInfoComponent},
  {path: 'profile/changeInfo/addSkill', component: AddSkillComponent},
  {path: 'profile/changeInfo/addExpirience', component: AddExpirienceComponent},
  {path: 'profile/changeInfo/addEducation', component: AddEducationComponent},
  {path: 'profile/changeInfo/addInterest', component: AddInterestComponent},

  {path: 'postFans', component: PostFansComponent},

  // prikazi profil korisnika
  {path: 'selectedUserProfile', component: SelectedUserProfileComponent},
  {path: 'selectedUserProfile/selectedUserPosts', component: SelectedUserPostsComponent},
  {path: 'selectedUserProfile/selectedUserFollowers', component: UserFollowersComponent},
  {path: 'selectedUserProfile/selectedFollowingUsers', component: FollowingUsersComponent},
  {path: 'selectedUserProfile/selectedUserFollowers', component: SelectedUserFollowersComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
