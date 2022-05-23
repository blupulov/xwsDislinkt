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

  constructor(private userService: UserService, private router: Router) { }
  
  ngOnInit(): void {
    
  }

  onSubmit() {
    this.userService.login(this.username, this.password).subscribe(
      res => {
        localStorage.setItem('userId', res.userId)
        localStorage.setItem('role', res.role)
        localStorage.setItem('token', res.token)
        this.router.navigateByUrl('/profile');
      },
      err => {
        alert("LOGIN PROBLEM");
      }
    )
  }

}
