import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Company } from 'src/models/company.model';
import { CompanyService } from 'src/services/company.service';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-show-my-companies',
  templateUrl: './show-my-companies.component.html',
  styleUrls: ['./show-my-companies.component.css']
})
export class ShowMyCompaniesComponent implements OnInit {
  companies: Company[] = [];
  
  constructor(private userService: UserService, private companyService: CompanyService, private router:Router) { }

  ngOnInit(): void {
    this.loadCompanies()
  }

  loadCompanies() {
    this.companyService.getOwnerCompanies(this.userService.getUserId()).subscribe(
      res => {
        this.companies = res.companies
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
