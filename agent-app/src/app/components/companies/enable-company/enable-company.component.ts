import { TOUCH_BUFFER_MS } from '@angular/cdk/a11y/input-modality/input-modality-detector';
import { Component, OnInit } from '@angular/core';
import { Company } from 'src/models/company.model';
import { CompanyService } from 'src/services/company.service';
import { CustomService } from 'src/services/custom.service';

@Component({
  selector: 'app-enable-company',
  templateUrl: './enable-company.component.html',
  styleUrls: ['./enable-company.component.css']
})
export class EnableCompanyComponent implements OnInit {
  companies: Company[] = [];
  
  constructor(private companyService: CompanyService, private customService: CustomService) { }

  ngOnInit(): void {
    this.loadCompanies()
  }

  loadCompanies() {
    this.companyService.getUnAcceptedCompanies().subscribe(
      res => {
        this.companies = res.companies
      }, err => {
        alert('problem with loading companies')
      }
    )
  }

  deleteCompany(id: string) {
    this.companyService.deleteCompany(id).subscribe(
      res => {
        this.loadCompanies()
      }, err => {
        alert('problem with deleting company')
      }
    )
  }

  enableCompany(userId: string, companyId: string) {
    this.customService.enableCompany(userId, companyId).subscribe(
      res => {
        this.loadCompanies()
      }, err => {
        alert('problem with enabling companies')
      }
    )
  }
}
