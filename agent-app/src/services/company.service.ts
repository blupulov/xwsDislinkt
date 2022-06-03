import { Injectable } from "@angular/core";
import { HttpClient } from '@angular/common/http';
import { UserService } from "./user.service";
import { CompaniesResponse, Company, Job, Comment } from "src/models/company.model";

@Injectable({
  providedIn: 'root'
})

export class CompanyService {

  constructor (private http: HttpClient, private userService: UserService) { }

  private apiUrl = 'http://localhost:8001/company';

  createCompany(company: Company) {
    return this.http.post(this.apiUrl, JSON.stringify(company))
  }

  getOwnerCompanies(ownerId: String | null) {
    return this.http.get<CompaniesResponse>(this.apiUrl + '/' + ownerId + '/owner')
  }

  getCompanyById() {
    return this.http.get<any>(this.apiUrl + '/' + this.getSelectedCompanyId())
  }

  createJob(job: Job) {
    return this.http.put(this.apiUrl + '/' + this.getSelectedCompanyId() + '/job', JSON.stringify(job))
  }

  deleteJob(jobId: string) {
    return this.http.delete(this.apiUrl + '/' + this.getSelectedCompanyId() + '/job/' + jobId)
  }

  createComment(comment: Comment) {
    return this.http.put(this.apiUrl + '/' + this.getSelectedCompanyId() + '/comment', JSON.stringify(comment))
  }

  deleteComment(commentId: string) {
    return this.http.delete(this.apiUrl + '/' + this.getSelectedCompanyId() + '/comment/' + commentId)
  }

  setSelectedCompanyId(companyId: string) {
    localStorage.setItem("selectedCompanyId", companyId)
  }

  getSelectedCompanyId():any {
    return localStorage.getItem("selectedCompanyId")
  }
}