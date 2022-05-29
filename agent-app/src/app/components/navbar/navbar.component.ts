import { Component, OnInit } from '@angular/core';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  constructor(public userService: UserService) { }

  ngOnInit(): void {
  }

  logout(): void {
    this.userService.logout()
  }

  onSubmit() {
    //datum treba prilagotiti
    this.userService.GetUserByUsername().subscribe(
      res => {
        alert("username is ")

      }, err => {
        alert('problem with sending comment')
      }
    )
  }
}
