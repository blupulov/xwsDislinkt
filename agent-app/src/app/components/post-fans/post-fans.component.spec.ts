import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PostFansComponent } from './post-fans.component';

describe('PostFansComponent', () => {
  let component: PostFansComponent;
  let fixture: ComponentFixture<PostFansComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PostFansComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PostFansComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
