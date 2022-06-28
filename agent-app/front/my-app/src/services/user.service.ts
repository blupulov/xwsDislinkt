import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { LoginResponse } from "src/models/loginResponse.model";
import { Registration } from "src/models/registration.model";

@Injectable({
  providedIn: 'root'
})

export class UserService {

  constructor (private http: HttpClient, private router: Router) { }

  private apiUrl = 'http://localhost:8053/agentApp/user';

  registration(newUser: Registration){
    let parts = newUser.birthDate.split('T')
    let dob = parts[0] + 'T00:00:00Z'
    newUser.birthDate = dob
    return this.http.post<any>(this.apiUrl + '/register', JSON.stringify(newUser));
  }

  login(username: String, password: String) {
    let params = {username: username, password: password}
    return this.http.post<LoginResponse>(this.apiUrl + '/login', JSON.stringify(params))
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

  getSelectedUserId() {
    return localStorage.getItem('selectedUserId') || '';
  }

  setSelectedUserId(userId: String) {
    localStorage.setItem('selectedUserId', userId.toString())
  }

}