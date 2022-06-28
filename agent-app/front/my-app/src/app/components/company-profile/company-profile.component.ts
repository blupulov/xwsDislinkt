import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Company } from 'src/models/company.model';
import { CompanyService } from 'src/services/company.service';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-company-profile',
  templateUrl: './company-profile.component.html',
  styleUrls: ['./company-profile.component.css']
})
export class CompanyProfileComponent implements OnInit {
  company: Company = new Company();
  
  constructor(private companyService: CompanyService, public userService: UserService, private router: Router) { }

  ngOnInit(): void {
    this.loadCompany()
  }

  loadCompany() {
    this.companyService.getCompanyById().subscribe(
      res => {
        this.company = res
      }, err => {
        alert('problem with loading company')
      }
    )
  }

  createJob() {
    this.router.navigateByUrl('createJob')
  }

  createComment() {
    this.router.navigateByUrl('createComment')
  }


  deleteJob(id: string) {
    this.companyService.deleteJob(id).subscribe(
      res => {
        this.loadCompany()
      }, err => {
        alert('problem with deleting job')
      }
    )
  }

  deleteComment(id: string) {
    this.companyService.deleteComment(id).subscribe(
      res => {
        this.loadCompany()
      }, err => {
        alert('problem with deleting comment')
      }
    )
  }
}
