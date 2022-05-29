import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-find-user',
  templateUrl: './find-user.component.html',
  styleUrls: ['./find-user.component.css']
})
export class FindUserComponent implements OnInit {

  username : String = '';
  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {

  }

  onSubmit() {
    //datum treba prilagotiti
    this.userService.GetUserByUsername(this.username).subscribe(
      res => {
          this.userService.setSelectedUserId(res.user.id)
          //alert(this.userService.getSelectedUserId())
          //this.router.navigateByUrl('selectedUserProfile')
      }, err => {
        alert('user not found')
      }
    )
  }
}
