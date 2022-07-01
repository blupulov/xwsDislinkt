import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-get-api-token',
  templateUrl: './get-api-token.component.html',
  styleUrls: ['./get-api-token.component.css']
})
export class GetApiTokenComponent implements OnInit {
  password: String = '';
  username: String = '';

  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {}

  onSubmit() {
    this.userService.createApiToken(this.username, this.password).subscribe(
      res => {
        localStorage.setItem('apiToken', res.token)
        this.router.navigateByUrl('companyProfile')
      }, err => {
        alert('problem with creating api token')
      }
    )
  }

}
