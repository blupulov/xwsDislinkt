import { Injectable } from "@angular/core";
import { HttpClient } from '@angular/common/http';
import { Login } from "src/models/login.model";
import { Router } from "@angular/router";
import { Registration } from "src/models/registration.model";
import { User, UserResponse } from "src/models/user.model";



@Injectable({
  providedIn: 'root'
})

export class UserService {

  constructor (private http: HttpClient, private router: Router) { }

  private apiUrl = 'http://localhost:8001/user';

  login(username: String, password: String) {
    let params = {username: username, password: password}
    return this.http.post<Login>(this.apiUrl + '/login', JSON.stringify(params));
  }

  registration(newUser: Registration){
    newUser.birthdate = "1999-11-11T00:00:00.000+00:00"
    return this.http.post<any>(this.apiUrl + '/register', JSON.stringify(newUser));
  }

  addSkill(newSkill: any) {
    return this.http.put(this.apiUrl + '/addSkill/' + this.getUserId(), JSON.stringify(newSkill))
  }

  getUserById(userId: String | null) {
    return this.http.get<UserResponse>(this.apiUrl + '/' + userId);
  }

  logout(): void {
    localStorage.clear()
    this.router.navigateByUrl('/login');
  }

  getUserRole(): String | null {
    return localStorage.getItem('role')
  }

  getToken(): String | null {
    return localStorage.getItem('token')
  }

  getUserId(): String | null {
    return localStorage.getItem('userId')
  }

}
