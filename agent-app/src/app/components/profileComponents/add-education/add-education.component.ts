import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AddEducation } from 'src/models/education.model';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-add-education',
  templateUrl: './add-education.component.html',
  styleUrls: ['./add-education.component.css']
})
export class AddEducationComponent implements OnInit {

  public newEducation: AddEducation = new AddEducation();

  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {
  }

  onSubmit(form: any){
    this.userService.AddEducation(this.newEducation).subscribe(
      res => {
        alert("education added")
        this.router.navigateByUrl("profile")
      },
      err => {
        alert("problem while adding education!")
      }
    )
}

}
