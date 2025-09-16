import { render, screen } from '@testing-library/svelte';
import Downloads from '../src/routes/downloads/+page.svelte';
import { describe, it, expect } from 'vitest';

describe('Downloads Page', () => {
  it('renders the heading', () => {
    render(Downloads);
    expect(screen.getByText('Downloads')).toBeInTheDocument();
  });
});