import { Injectable } from "@angular/core";
import { HttpClient } from '@angular/common/http';
import { UserService } from "./user.service";
import { UserIds } from "src/models/user.model";

@Injectable({
  providedIn: 'root'
})

export class FollowingService {

  constructor(private http: HttpClient, private userService: UserService) {}
  
  private apiUrl = 'http://localhost:8001/follow';

  getAllFollowingUsersIds(userId: String | null) {
    return this.http.get<UserIds>(this.apiUrl + '/' + userId + '/following')
  }

  getAllUserFollowersIds(userId: String | null) {
    return this.http.get<UserIds>(this.apiUrl + '/' + userId + '/followers')
  }

  // zapracivanje i otpracivanje je moguce izvest samo sa profila selektovanog usera pa se podaci vuku iz ls
  follow() {
    return this.http.post(this.apiUrl + '/' + this.userService.getUserId() + '/follow/'
      + this.userService.getSelectedUserId(), {})
  }
  unFollow() {
    return this.http.delete(this.apiUrl + '/' + this.userService.getUserId() + '/unfollow/'
      + this.userService.getSelectedUserId())
  }

}