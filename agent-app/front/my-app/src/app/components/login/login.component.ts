import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  password: String = '';
  username: String = '';

  constructor(private router: Router, private userService: UserService) { }

  ngOnInit(): void { }

  onSubmit() {
    this.userService.login(this.username, this.password).subscribe(
      res => {
        localStorage.setItem('userId', res.userId)
        localStorage.setItem('role', res.role)
        localStorage.setItem('token', res.token)
        this.router.navigateByUrl('/userProfile');
      },
      err => {
        alert("LOGIN PROBLEM");
      }
    )
  }


}
