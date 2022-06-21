import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Company } from 'src/models/company.model';
import { CompanyService } from 'src/services/company.service';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-show-user-companies',
  templateUrl: './show-user-companies.component.html',
  styleUrls: ['./show-user-companies.component.css']
})
export class ShowUserCompaniesComponent implements OnInit {
  companies: Company[] = [];
  
  constructor(private userService: UserService, private router: Router, private companyService: CompanyService) { }

  ngOnInit(): void {
    this.loadCompanies()
  }

  loadCompanies() {
    this.companyService.getOwnerCompanies(this.userService.getSelectedUserId()).subscribe(
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
