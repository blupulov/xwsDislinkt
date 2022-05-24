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
import { CommentPostComponent } from './components/profileComponents/comment-post/comment-post.component';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatDividerModule} from '@angular/material/divider';
import {MatTableModule} from '@angular/material/table';

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
    CommentPostComponent,
  ],
  imports: [
    BrowserModule,
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
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
