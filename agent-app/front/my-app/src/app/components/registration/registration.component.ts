import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Registration } from 'src/models/registration.model';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent implements OnInit {
  url = "./assets/no-image.png"
  newUser: Registration = new Registration();

  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void { }

  onSelectFile(image: any) {
    if(image.target.files) {
      var reader = new FileReader();
      reader.readAsDataURL(image.target.files[0]);
      reader.onload = (event: any) => {
        this.url = event.target.result;
        this.newUser.profileImage = event.target.result;
      }
    }
  }

  onSubmit(form: any){
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
