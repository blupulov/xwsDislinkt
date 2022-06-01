import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-mid-comp',
  templateUrl: './mid-comp.component.html',
  styleUrls: ['./mid-comp.component.css']
})
export class MidCompComponent implements OnInit {

  constructor(private router: Router) { }

  ngOnInit(): void {
    this.router.navigateByUrl('selectedUserProfile')
  }

}
