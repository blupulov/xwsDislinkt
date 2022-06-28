import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { Company } from "src/models/company.model";


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

  getCompanyById(companyId: String | null) {
    return this.http.get<Company>(this.apiUrl + '/companyId/' + companyId)
  }

  getAllCompaniesByOwnerId(ownerId: String | null) {
    return this.http.get<Company[]>(this.apiUrl + '/ownerId/' + ownerId)
  }

  setSelectedCompanyId(companyId: string) {
    localStorage.setItem("selectedCompanyId", companyId)
  }

  getSelectedCompanyId():any {
    return localStorage.getItem("selectedCompanyId")
  }

}