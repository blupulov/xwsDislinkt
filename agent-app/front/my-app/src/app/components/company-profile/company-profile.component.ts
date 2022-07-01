import { ThrowStmt } from '@angular/compiler';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Company, Job } from 'src/models/company.model';
import { User } from 'src/models/user.model';
import { CompanyService } from 'src/services/company.service';
import { CustomService } from 'src/services/custom.service';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-company-profile',
  templateUrl: './company-profile.component.html',
  styleUrls: ['./company-profile.component.css']
})
export class CompanyProfileComponent implements OnInit {
  company: Company = new Company();
  companyOwner: User = new User();

  constructor(private companyService: CompanyService, public userService: UserService, private router: Router, 
    private customService: CustomService) { }

  ngOnInit(): void {
    this.loadCompany()
  }

  loadCompany() {
    this.companyService.getCompanyById().subscribe(
      res => {
        this.company = res
        this.loadCompanyOwner(res.ownerId)
      }, err => {
        alert('problem with loading company')
      }
    )
  }

  loadCompanyOwner(userId: String) {
    this.userService.getUserById(userId).subscribe(
      res => {
        this.companyOwner = res
      }, err => {
        alert('problem with loading company owner')
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

  shareJobOnDislinkt(job: Job) {
    if (this.userService.getApiToken() == ''){
      this.router.navigateByUrl('getApiToken')
    } else {
      let token = this.userService.getApiToken()
      let username = this.companyOwner.username.toString()
      this.customService.shareJobOnDislinkt(job, token, this.company.companyName, username).subscribe(
        res => {
          alert('job is shared')
        }, err => {
          alert('problem with sharing job')
        }
      )
    }
  }
}
