import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { Job } from "src/models/company.model";

@Injectable({
  providedIn: 'root'
})

export class CustomService {

  constructor (private http: HttpClient, private router: Router) { }

  private apiUrl = 'http://localhost:8053/agentApp';

  enableCompany(userId: string, companyId: string) {
    return this.http.put(this.apiUrl + '/enable/' + companyId + '/promote/' + userId, {})
  }

  shareJobOnDislinkt(job: Job, token: string, companyName: string, username: string ) {
    let params  = {apiToken: token,
      jobName: job.jobName,
      companyName: companyName,
      disJobUsername: username}
    return this.http.post('http://localhost:8001/post/job', JSON.stringify(params))
  }

}