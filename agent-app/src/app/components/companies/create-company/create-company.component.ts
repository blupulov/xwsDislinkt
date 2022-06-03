import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Company } from 'src/models/company.model';
import { CompanyService } from 'src/services/company.service';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-create-company',
  templateUrl: './create-company.component.html',
  styleUrls: ['./create-company.component.css']
})
export class CreateCompanyComponent implements OnInit {

  newCompany: Company = new Company();

  constructor(private userService: UserService, private companyService: CompanyService, private router: Router) { }

  ngOnInit(): void {

  }

  onSubmit() {
    this.newCompany.ownerId = this.userService.getUserId()?.toString() || '';
    this.newCompany.info.phoneNumber = this.newCompany.info.phoneNumber.toString();
    this.companyService.createCompany(this.newCompany).subscribe(
      res => {
        alert('company request created')
        this.router.navigateByUrl('profile')
      }, err => {
        alert('problem with creating company')
      }
    )
  }

}
