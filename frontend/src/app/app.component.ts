import { Component, DestroyRef, inject, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { BackendService } from './services/backend/backend.service';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
})
export class AppComponent implements OnInit {
  private backendService = inject(BackendService);
  private destroyRef = inject(DestroyRef);

  title = 'frontend';
  m = this.backendService.loadedMessage;

  ngOnInit(): void {
    const subscription = this.backendService.loadMessage().subscribe({
      error: (error) => {
        console.log('error: ', error.message);
      },
      complete: () => {
        console.log('done fetching');
      },
    });

    this.destroyRef.onDestroy(() => {
      subscription.unsubscribe();
    });
  }
}
