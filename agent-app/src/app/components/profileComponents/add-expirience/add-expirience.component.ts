import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AddExpirience } from 'src/models/expirience.model';
import { UserService } from 'src/services/user.service';
import {MatFormFieldModule} from '@angular/material/form-field';

@Component({
  selector: 'app-add-expirience-component',
  templateUrl: './add-expirience.component.html',
  styleUrls: ['./add-expirience.component.css']
})
export class AddExpirienceComponent implements OnInit {

  public newExpirience: AddExpirience = new AddExpirience();


  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {
  }

  onSubmit(form: any){
    this.userService.addExpirience(this.newExpirience).subscribe(
      res => {
        alert("expirience added")
        this.router.navigateByUrl("profile")
      },
      err => {
        alert("problem while adding expirience!")
      }
    )
}

}
