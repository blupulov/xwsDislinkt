import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ToolbarComponent } from './components/toolbar/toolbar.component';
import { LoginComponent } from './components/login/login.component';
import { RegistrationComponent } from './components/registration/registration.component';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { UserProfileComponent } from './components/user-profile/user-profile.component';
import { FlexLayoutModule } from '@angular/flex-layout';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatCardModule } from '@angular/material/card';
import { FormsModule } from "@angular/forms";
import { ReactiveFormsModule } from "@angular/forms";
import { MatGridListModule } from '@angular/material/grid-list';
import { MatDividerModule } from '@angular/material/divider';
import { CreateCompanyComponent } from './components/create-company/create-company.component';
import { ChangeUserComponent } from './components/change-user/change-user.component';
import { EnableCompanyComponent } from './components/enable-company/enable-company.component';
import { CompanyProfileComponent } from './components/company-profile/company-profile.component';
import { SelectedUserCompaniesComponent } from './components/selected-user-companies/selected-user-companies.component';
import { MyCompaniesComponent } from './components/my-companies/my-companies.component';
import { CreateJobComponent } from './components/create-job/create-job.component';
import { CreateCommentComponent } from './components/create-comment/create-comment.component';
import { MatTabsModule } from '@angular/material/tabs';
import { SelectedUserProfileComponent } from './components/selected-user-profile/selected-user-profile.component';
import { MidSearchComponent } from './components/mid-search/mid-search.component';
import { FindUserComponent } from './components/find-user/find-user.component';
import { GetApiTokenComponent } from './components/get-api-token/get-api-token.component';

@NgModule({
  declarations: [
    AppComponent,
    ToolbarComponent,
    LoginComponent,
    RegistrationComponent,
    UserProfileComponent,
    CreateCompanyComponent,
    ChangeUserComponent,
    EnableCompanyComponent,
    CompanyProfileComponent,
    SelectedUserCompaniesComponent,
    MyCompaniesComponent,
    CreateJobComponent,
    CreateCommentComponent,
    SelectedUserProfileComponent,
    MidSearchComponent,
    FindUserComponent,
    GetApiTokenComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatToolbarModule,
    HttpClientModule,
    MatButtonModule,
    FlexLayoutModule,
    MatFormFieldModule,
    MatInputModule,
    MatCardModule,
    FormsModule,
    ReactiveFormsModule,
    MatGridListModule,
    MatDividerModule,
    MatTabsModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
