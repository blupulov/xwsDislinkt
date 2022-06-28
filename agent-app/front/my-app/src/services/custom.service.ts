import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";

@Injectable({
  providedIn: 'root'
})

export class CustomService {

  constructor (private http: HttpClient, private router: Router) { }

  private apiUrl = 'http://localhost:8053/agentApp';

  enableCompany(userId: string, companyId: string) {
    return this.http.put(this.apiUrl + '/enable/' + companyId + '/promote/' + userId, {})
  }
  
}