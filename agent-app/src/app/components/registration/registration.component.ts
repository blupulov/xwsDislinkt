import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Registration } from 'src/models/registration.model';
import { UserService } from 'src/services/user.service';
import {MatFormFieldModule} from '@angular/material/form-field';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent implements OnInit {
  password: String = '';
  username: String = '';
  biography: String = '';
  firstname: String = '';
  lastname: String = '';
  birthdate: String = '';
  email: String = '';
  phoneNumber:String = '';



  public newUser: Registration = new Registration();

  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {
  }

  onSubmit(form: any){
    console.log(this.newUser)
    if(this.newUser.password === this.newUser.confirmPassword){
      this.userService.registration(this.newUser).subscribe(
        res => {
          alert("registration complete")
          this.router.navigateByUrl("login")
        },
        err => {
          alert("problem with registration!")
        }
      )
    }
    else{
      alert("Password and confirmed password must be the same!")
    }
  }

}
