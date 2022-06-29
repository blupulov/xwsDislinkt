import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Job } from 'src/models/company.model';
import { CompanyService } from 'src/services/company.service';

@Component({
  selector: 'app-create-job',
  templateUrl: './create-job.component.html',
  styleUrls: ['./create-job.component.css']
})
export class CreateJobComponent implements OnInit {
  job: Job = new Job();

  constructor(private companyService: CompanyService, private router: Router) { }

  ngOnInit(): void { }

  onSubmit() {
    this.companyService.createJob(this.job).subscribe(
      res => {
        this.router.navigateByUrl('companyProfile')
      }, err => {
        alert('problem with creating job')
      }
    )
  }
} 
