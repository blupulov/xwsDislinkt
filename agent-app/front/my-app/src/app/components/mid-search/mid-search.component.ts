import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-mid-search',
  templateUrl: './mid-search.component.html',
  styleUrls: ['./mid-search.component.css']
})
export class MidSearchComponent implements OnInit {

  constructor(private router: Router) { }

  ngOnInit(): void {
    this.router.navigateByUrl('selectedUserProfile')
  }

}
