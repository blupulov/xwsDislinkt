import { Injectable } from "@angular/core";
import { HttpClient } from '@angular/common/http';
import { Login } from "src/models/login.model";
import { Router } from "@angular/router";

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