import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Company } from 'src/models/company.model';
import { CompanyService } from 'src/services/company.service';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-selected-user-companies',
  templateUrl: './selected-user-companies.component.html',
  styleUrls: ['./selected-user-companies.component.css']
})
export class SelectedUserCompaniesComponent implements OnInit {
  companies: Company[] = [];
  
  constructor(private userService: UserService, private router: Router, private companyService: CompanyService) { }

  ngOnInit(): void {
    this.loadCompanies()
  }

  loadCompanies() {
    this.companyService.getAllCompaniesByOwnerId(this.userService.getSelectedUserId()).subscribe(
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
