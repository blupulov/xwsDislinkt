import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { LoginComponent } from './components/login/login.component';
import { RegistrationComponent } from './components/registration/registration.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatButtonModule } from '@angular/material/button';
import { FlexLayoutModule } from '@angular/flex-layout';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatCardModule } from '@angular/material/card';
import { FormsModule } from "@angular/forms";
import { ReactiveFormsModule } from "@angular/forms";
import { MatToolbarModule } from "@angular/material/toolbar";
import { ProfileComponent } from './components/profileComponents/profile/profile.component';
import { UserPostsComponent } from './components/profileComponents/user-posts/user-posts.component';
import { CreatePostComponent } from './components/profileComponents/create-post/create-post.component';
import { ChangeInfoComponent } from './components/profileComponents/change-info/change-info.component';
import { UserFollowersComponent } from './components/profileComponents/user-followers/user-followers.component';
import { FollowingUsersComponent } from './components/profileComponents/following-users/following-users.component';
import { MainPageComponent } from './components/main-page/main-page.component';
import { SelectedUserProfileComponent } from './components/selected-user-profile/selected-user-profile.component';
import { HttpClientModule } from '@angular/common/http';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatDividerModule} from '@angular/material/divider';
import {MatTableModule} from '@angular/material/table';
import { AddSkillComponent } from './components/profileComponents/add-skill/add-skill.component';
import { FindUserComponent } from './components/find-user/find-user.component';
import { AddExpirienceComponent } from './components/profileComponents/add-expirience/add-expirience.component';
import { AddEducationComponent } from './components/profileComponents/add-education/add-education.component';
import { AddInterestComponent } from './components/profileComponents/add-interest/add-interest.component';
import { PostHatersComponent } from './components/post-haters/post-haters.component';
import { PostFansComponent } from './components/post-fans/post-fans.component';
import {MatDatepickerModule} from '@angular/material/datepicker';
import { SelectedUserPostsComponent } from './components/selected-user-posts/selected-user-posts.component';
import { SelectedUserFollowersComponent } from './components/selected-user-followers/selected-user-followers.component';
import { MidCompComponent } from './mid-comp/mid-comp.component';
import { SelectedUserFollowingComponent } from './components/selected-user-following/selected-user-following.component';
import { CreateCompanyComponent } from './components/companies/create-company/create-company.component';
import { ShowUserCompaniesComponent } from './components/companies/show-user-companies/show-user-companies.component';
import { CompanyProfileComponent } from './components/companies/company-profile/company-profile.component';
import { EnableCompanyComponent } from './components/companies/enable-company/enable-company.component';
import { ShowMyCompaniesComponent } from './components/companies/show-my-companies/show-my-companies.component';
import { CreateJobComponent } from './components/companies/create-job/create-job.component';
import { CreateCommentComponent } from './components/companies/create-comment/create-comment.component';

@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    LoginComponent,
    RegistrationComponent,
    ProfileComponent,
    UserPostsComponent,
    CreatePostComponent,
    ChangeInfoComponent,
    UserFollowersComponent,
    FollowingUsersComponent,
    MainPageComponent,
    SelectedUserProfileComponent,
    AddSkillComponent,
    FindUserComponent,
    AddExpirienceComponent,
    AddEducationComponent,
    PostHatersComponent,
    PostFansComponent,
    AddInterestComponent,
    SelectedUserPostsComponent,
    SelectedUserFollowersComponent,
    MidCompComponent,
    SelectedUserFollowingComponent,
    CreateCompanyComponent,
    ShowUserCompaniesComponent,
    CompanyProfileComponent,
    EnableCompanyComponent,
    ShowMyCompaniesComponent,
    CreateJobComponent,
    CreateCommentComponent,

  ],
  imports: [
    AppRoutingModule,
    BrowserAnimationsModule,
    MatButtonModule,
    FlexLayoutModule,
    MatFormFieldModule,
    MatInputModule,
    MatCardModule,
    FormsModule,
    ReactiveFormsModule,
    MatToolbarModule,
    HttpClientModule,
    MatGridListModule,
    MatDividerModule,
    MatTableModule,
    MatDatepickerModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
