import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { Company, Job, Comment } from "src/models/company.model";


@Injectable({
  providedIn: 'root'
})

export class CompanyService {

  constructor (private http: HttpClient, private router: Router) { }

  private apiUrl = 'http://localhost:8053/agentApp/company';

  createCompany(company: Company) {
    return this.http.post(this.apiUrl, JSON.stringify(company))
  }

  getUnAcceptedCompanies() {
    return this.http.get<Company[]>(this.apiUrl + '/unAccepted')
  }

  getCompanyById() {
    return this.http.get<Company>(this.apiUrl + '/companyId/' + this.getSelectedCompanyId())
  }

  getAllCompaniesByOwnerId(ownerId: String | null) {
    return this.http.get<Company[]>(this.apiUrl + '/ownerId/' + ownerId)
  }

  deleteCompany(companyId: string) {
    return this.http.delete(this.apiUrl + '/' + companyId)
  }

  setSelectedCompanyId(companyId: string) {
    localStorage.setItem("selectedCompanyId", companyId)
  }

  getSelectedCompanyId():any {
    return localStorage.getItem("selectedCompanyId")
  }

  createJob(job: Job) {
    return this.http.put(this.apiUrl + '/job/' + this.getSelectedCompanyId(), JSON.stringify(job))
  }

  deleteJob(jobId: string) {
    return this.http.delete(this.apiUrl + '/' + this.getSelectedCompanyId() + '/job/' + jobId)
  }

  createComment(comment: Comment) {
    return this.http.put(this.apiUrl + '/comment/' + this.getSelectedCompanyId(), JSON.stringify(comment))
  }

  deleteComment(commentId: string) {
    return this.http.delete(this.apiUrl + '/' + this.getSelectedCompanyId() + '/comment/' + commentId)
  }
}