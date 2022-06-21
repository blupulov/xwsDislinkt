import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AddSkill } from 'src/models/skill.model';
import { UserService } from 'src/services/user.service';
import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';


@Component({
  selector: 'app-add-skill',
  templateUrl: './add-skill.component.html',
  styleUrls: ['./add-skill.component.css']
})
export class AddSkillComponent implements OnInit {

  public newSkill: AddSkill = new AddSkill();

  constructor(private userService: UserService, private router: Router) { }


  ngOnInit(): void {
  }

  onSubmit(form: any){
      this.userService.addSkill(this.newSkill).subscribe(
        res => {
          alert("skill added")
          this.router.navigateByUrl("profile")
        },
        err => {
          alert("problem while adding skill!")
        }
      )
  }
}
