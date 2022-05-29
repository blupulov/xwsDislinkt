import { Injectable } from "@angular/core";
import { HttpClient } from '@angular/common/http';
import { Router } from "@angular/router";
import { User } from "src/models/user.model";

@Injectable({
  providedIn: 'root'
})

export class CustomService {

  constructor (private http: HttpClient, private router: Router) { }

  private apiUrl = 'http://localhost:8001/custom';

  //oni sto ga prate
  getUserFollowers(userId: string | null | undefined) {
    return this.http.get<User[]>(this.apiUrl + '/user/' + userId + "/followers")
  }
  // oni koje prati
  getFollowingUsers(userId: string | null | undefined) {
    return this.http.get<User[]>(this.apiUrl + '/user/' + userId + "/following")
  }

  // samo id-jevi osoba koje ulogovani korisnik prati
  setFollowing(usersIds: String[]) {
    localStorage.setItem("following", JSON.stringify(usersIds))
  }

  getFollowing() {
    return JSON.parse(localStorage.getItem("following") || "{}")
  }

  // nakon svakog uspesnog zapracivanja ubaci se u listu novi userId
  pushFollowing(userId: string) {
    let following = this.getFollowing()
    following.push(userId)
    this.setFollowing(following)
  }
  
}