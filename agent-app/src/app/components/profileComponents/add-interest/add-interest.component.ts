import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AddInterest } from 'src/models/interest.model';
import { UserService } from 'src/services/user.service';

@Component({
  selector: 'app-add-interest',
  templateUrl: './add-interest.component.html',
  styleUrls: ['./add-interest.component.css']
})
export class AddInterestComponent implements OnInit {

  public newInterest: AddInterest = new AddInterest();

  constructor(private userService: UserService, private router: Router) { }


  ngOnInit(): void {
  }


  onSubmit(form: any){
    this.userService.AddInterest(this.newInterest).subscribe(
      res => {
        alert("interest added")
        this.router.navigateByUrl("profile")
      },
      err => {
        alert("problem while adding interest!")
      }
    )
}

}
