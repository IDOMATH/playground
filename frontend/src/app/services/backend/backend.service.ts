import { HttpClient } from '@angular/common/http';
import { inject, Injectable, signal } from '@angular/core';
import { catchError, map, tap, throwError } from 'rxjs';

interface Message {
  msg: string;
}

@Injectable({
  providedIn: 'root',
})
export class BackendService {
  private httpClient = inject(HttpClient);
  private m = signal<string>('');

  loadedMessage = this.m.asReadonly();

  loadMessage() {
    return this.fetchMessage(
      'http://localhost:8080/',
      'something went wrong'
    ).pipe(
      tap({
        next: (m) => this.m.set(m),
      })
    );
  }

  private fetchMessage(url: string, errorMessage: string) {
    return this.httpClient.get<{ msg: string }>(url).pipe(
      map((resData) => resData.msg),
      catchError((error) => {
        console.log(error);
        return throwError(() => new Error(errorMessage));
      })
    );
  }
}
