import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PostHatersComponent } from './post-haters.component';

describe('PostHatersComponent', () => {
  let component: PostHatersComponent;
  let fixture: ComponentFixture<PostHatersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PostHatersComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PostHatersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
