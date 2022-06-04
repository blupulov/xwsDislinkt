import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CompanyProfileComponent } from './components/companies/company-profile/company-profile.component';
import { CreateCommentComponent } from './components/companies/create-comment/create-comment.component';
import { CreateCompanyComponent } from './components/companies/create-company/create-company.component';
import { CreateJobComponent } from './components/companies/create-job/create-job.component';
import { EnableCompanyComponent } from './components/companies/enable-company/enable-company.component';
import { ShowMyCompaniesComponent } from './components/companies/show-my-companies/show-my-companies.component';
import { ShowUserCompaniesComponent } from './components/companies/show-user-companies/show-user-companies.component';
import { LoginComponent } from './components/login/login.component';
import { MainPageComponent } from './components/main-page/main-page.component';
import { PostFansComponent } from './components/post-fans/post-fans.component';
import { PostHatersComponent } from './components/post-haters/post-haters.component';
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
import { SelectedUserFollowingComponent } from './components/selected-user-following/selected-user-following.component';
import { SelectedUserPostsComponent } from './components/selected-user-posts/selected-user-posts.component';
import { SelectedUserProfileComponent } from './components/selected-user-profile/selected-user-profile.component';
import { MidCompComponent } from './mid-comp/mid-comp.component';

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
  {path: 'profile/createCompany', component: CreateCompanyComponent},
  {path: 'profile/myCompanies', component: ShowMyCompaniesComponent},
  
  {path: 'postFans', component: PostFansComponent},
  {path: 'postHaters', component: PostHatersComponent},
  {path: 'companyProfile', component: CompanyProfileComponent},
  {path: 'enableCompany', component: EnableCompanyComponent},
  {path: 'createJob', component: CreateJobComponent},
  {path: 'createComment', component: CreateCommentComponent},

  //KOMPONENTE KOJIMA SE PRISTUPA SA SELEKTOVANOG PROFILA
  {path: 'selectedUserProfile', component: SelectedUserProfileComponent},
  {path: 'selectedUserProfile/selectedUserPosts', component: SelectedUserPostsComponent},
  {path: 'selectedUserProfile/selectedUserFollowers', component: SelectedUserFollowersComponent},
  {path: 'selectedUserProfile/selectedUserFollowing', component: SelectedUserFollowingComponent},
  {path: 'selectedUserProfile/userCompanies', component: ShowUserCompaniesComponent},

  {path: 'mid', component: MidCompComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
