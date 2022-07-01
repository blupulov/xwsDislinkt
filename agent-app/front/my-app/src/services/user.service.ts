import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { LoginResponse } from "src/models/loginResponse.model";
import { Registration } from "src/models/registration.model";
import { User, UserResponse } from "src/models/user.model";

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

  getUserById(userId: String | null) {
    return this.http.get<User>(this.apiUrl + '/' + userId)
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

  getUserByUsername(username: String) {
    return this.http.get<User>(this.apiUrl + '/0/' + '/' + username)
  }

  getApiToken() {
    return localStorage.getItem('apiToken') || ''
  }

  createApiToken(username: String, password: String) {
    let params = {username: username, password: password}
    return this.http.post<any>('http://localhost:8001/user/apiToken', JSON.stringify(params))
  }

}