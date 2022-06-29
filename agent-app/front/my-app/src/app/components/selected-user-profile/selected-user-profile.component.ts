import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-selected-user-profile',
  templateUrl: './selected-user-profile.component.html',
  styleUrls: ['./selected-user-profile.component.css']
})
export class SelectedUserProfileComponent implements OnInit {

  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {
    if(this.userService.getSelectedUserId() == this.userService.getUserId())
      this.router.navigateByUrl('userProfile')
  }

}
