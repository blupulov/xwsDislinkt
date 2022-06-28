import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Company } from 'src/models/company.model';
import { CompanyService } from 'src/services/company.service';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-my-companies',
  templateUrl: './my-companies.component.html',
  styleUrls: ['./my-companies.component.css']
})
export class MyCompaniesComponent implements OnInit {
  companies: Company[] = [];
  
  constructor(private userService: UserService, private companyService: CompanyService, private router:Router) { }

  ngOnInit(): void {
    this.loadCompanies()
  }

  loadCompanies() {
    this.companyService.getAllCompaniesByOwnerId(this.userService.getUserId()).subscribe(
      res => {
        this.companies = res
      }, err => {
        alert('problem with loading companies')
      }
    )
  }

  showCompany(companyId: string) {
    this.companyService.setSelectedCompanyId(companyId)
    this.router.navigateByUrl('companyProfile')
  }
}
